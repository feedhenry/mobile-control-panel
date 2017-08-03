/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package start

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"os"

	"time"

	"github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile/v1alpha1"
	"github.com/feedhenry/mobile-control-panel/server/pkg/apiserver"
	clientset "github.com/feedhenry/mobile-control-panel/server/pkg/client/clientset_generated/internalclientset"
	"github.com/pborman/uuid"
	"k8s.io/apiserver/pkg/authorization/authorizerfactory"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

const defaultEtcdPathPrefix = "/registry/mobile.kubernetes.io"

// MobileServerOptions are the options used when starting the mobile apiserver
type MobileServerOptions struct {
	// the runtime configuration of our server
	GenericServerRunOptions *genericoptions.ServerRunOptions
	// the https configuration. certs, etc
	//ServingOptions *genericserveroptions.ServingOptions
	ServingOptions *genericoptions.SecureServingOptions
	// storage with etcd
	EtcdOptions *genericoptions.EtcdOptions
	// authn
	AuthenticationOptions *genericoptions.DelegatingAuthenticationOptions
	// authz
	AuthorizationOptions *genericoptions.DelegatingAuthorizationOptions

	RecommendedOptions *genericoptions.RecommendedOptions
	// StandaloneMode is used when running outside of kubernetes
	StandaloneMode bool
}

func NewMobileServerOptions() *MobileServerOptions {
	recommended := genericoptions.NewRecommendedOptions(defaultEtcdPathPrefix, apiserver.Scheme, apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion))
	o := &MobileServerOptions{
		RecommendedOptions:      recommended,
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		ServingOptions:          genericoptions.NewSecureServingOptions(),
		EtcdOptions:             recommended.Etcd,
		AuthenticationOptions:   recommended.Authentication,
		AuthorizationOptions:    recommended.Authorization,
	}
	fmt.Println("mobile options request headers", o.AuthenticationOptions.RequestHeader.UsernameHeaders)
	return o
}

func (o *MobileServerOptions) addFlags(flags *pflag.FlagSet) {
	o.RecommendedOptions.AddFlags(flags)
}

// NewCommandStartMobileServer provides a CLI handler for 'start master' command
func NewCommandStartMobileServer(stopCh <-chan struct{}) *cobra.Command {
	o := NewMobileServerOptions()

	cmd := &cobra.Command{
		Short: "Launch a mobile API server",
		Long:  "Launch a mobile API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunMobileServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.addFlags(flags)

	return cmd
}

func (o MobileServerOptions) Validate(args []string) error {
	return nil
}

func (o *MobileServerOptions) Complete() error {
	return nil
}

func (o *MobileServerOptions) standAlone() bool {
	return "true" == os.Getenv("MOBILE_API_SERVER_STANDALONE")
}

// TODO LOOK CLOSELY AT https://github.com/kubernetes-incubator/service-catalog/blob/4679685a7364cd3f8dd999d7205654589e19fbdb/cmd/apiserver/app/server/util.go#L52

func (o MobileServerOptions) Config() (*apiserver.Config, error) {
	// TODO have a "real" external address
	serverConfig := genericapiserver.NewConfig()
	serverConfig.WithSerializer(apiserver.Codecs)
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", net.ParseIP("127.0.0.1")); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}
	if !o.standAlone() {

		fmt.Println("server starting in cluster mode")
		if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
			return nil, errors.Wrap(err, "failed to applyto serverconfig")
		}

	} else {
		//allow for local dev without kubernetes
		o.RecommendedOptions.SecureServing.ServingOptions.BindPort = 3101
		o.RecommendedOptions.Authentication.SkipInClusterLookup = true
		o.RecommendedOptions.SecureServing.ServingOptions.BindAddress = net.ParseIP("0.0.0.0")
		etcdURL, ok := os.LookupEnv("KUBE_INTEGRATION_ETCD_URL")
		if !ok {
			etcdURL = "http://127.0.0.1:2379"
		}
		o.RecommendedOptions.Etcd.StorageConfig.ServerList = []string{etcdURL}
		o.RecommendedOptions.Etcd.StorageConfig.Prefix = uuid.New()
		serverConfig.Authenticator = nil
		serverConfig.AuditWriter = os.Stdout
		serverConfig.Authorizer = authorizerfactory.NewAlwaysAllowAuthorizer()
		if err := o.RecommendedOptions.Etcd.ApplyTo(serverConfig); err != nil {
			return nil, err
		}
		if err := o.RecommendedOptions.SecureServing.ApplyTo(serverConfig); err != nil {
			return nil, err
		}
		if err := o.RecommendedOptions.Audit.ApplyTo(serverConfig); err != nil {
			return nil, err
		}
		if err := o.RecommendedOptions.Features.ApplyTo(serverConfig); err != nil {
			return nil, err
		}
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
	}
	return config, nil
}

func (o MobileServerOptions) RunMobileServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return errors.Wrap(err, "failed to get serverconfig")
	}

	server, err := config.Complete().New()
	if err != nil {
		return errors.Wrap(err, "failed to Complete config")
	}
	time.AfterFunc(time.Second*2, func() {
		client, err := clientset.NewForConfig(server.GenericAPIServer.LoopbackClientConfig)
		b, err := client.Discovery().RESTClient().Get().AbsPath("/apis/mobile.k8s.io/v1alpha1").DoRaw()
		fmt.Println("discovery", err, string(b))
	})

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
