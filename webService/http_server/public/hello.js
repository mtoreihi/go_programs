angular.module('demo', [])
.controller('Hello', function($scope, $http) {
    $http.get('https://127.0.0.1:8443/getCount2').
        then(function(response) {
            //$scope.getCount2 = response.data;
            $scope.greeting = [{"id":123456,"content":"Hello, World!"}]
        });
});