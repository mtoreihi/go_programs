angular.module('demo', [])
.controller('Hello', function($scope, $http) {
	//$scope.greeting = [{"id":123456,"content":"Hello, World!"}]
    $http.get('http://127.0.0.1:8443/getCount2').
        then(function(response) {
            $scope.greeting = response.data;
            //$scope.greeting = [{"id":123456,"content":"Hello, World!"}]
        })
        .catch(function (err) {});
});