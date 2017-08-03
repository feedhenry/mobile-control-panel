console.log('Mobile Control Panel Extension Loaded');

window.OPENSHIFT_CONSTANTS.PROJECT_NAVIGATION.unshift({
  label: "Mobile",
  iconClass: "fa fa-mobile",
  href: "/mobile",
  prefixes: ["/mobile/"],
  isValid: function() {
    // TODO: Can this check if any mobile apps exist first?
    return true;
  }
});

angular
  .module('mobileOverviewExtension', ['openshiftConsole'])
  .config(function($routeProvider) {
    $routeProvider
      .when('/project/:project/mobile', {
        templateUrl: ' extensions/mcp/mobile.html',
        controller: 'MobileOverviewController'
      });
    }
  )
  .controller('MobileOverviewController', ['$scope', '$controller', '$routeParams', 'ProjectsService', 'APIService', 'DataService', function ($scope, $controller, $routeParams, ProjectsService, APIService, DataService) {
    // Initialize the super class and extend it.
    angular.extend(this, $controller('OverviewController', {$scope: $scope}));
    console.log('MobileOverviewController');

    ProjectsService
      .get($routeParams.project)
      .then(_.spread(function(project, context) {
        
        $scope.project = project;
        $scope.context = context;
        DataService.list({
            group: 'mobile.k8s.io',
            resource: 'mobileapps'
          }, $scope.context).then(function(resources) {
          $scope.mobileapps = resources.by("metadata.name");
          $scope.emptyMessage = "No " + APIService.kindToResource('MobileApp', true) + " to show";
        });

      }));
   }]);

hawtioPluginLoader.addModule('mobileOverviewExtension');