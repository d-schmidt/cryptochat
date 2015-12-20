var chatApp = angular.module('chatApp', ['luegg.directives']);

chatApp.factory('MyService', ['$q', '$rootScope', function($q, $rootScope) {
    var Service = {};
    var callbacks = {};
    var currentCallbackId = 0;
    var ws = new WebSocket("ws://127.0.0.1:12345/ws?encoding=text");

    ws.onopen = function(){
        console.log("Socket has been opened!");
    };

    ws.onclose = function(event){
        console.log("Socket has been closed!", event);
        $rootScope.$broadcast('close')
    };

    ws.onmessage = function(message) {
        listener(JSON.parse(message.data));
    };

    function waitForSocketConnection(socket, callback){
        setTimeout(
            function () {
                if (socket.readyState === 1) {
                    console.log("Connection is made")
                    if(callback != null){
                        callback();
                    }
                    return;
                } else {
                    console.log("wait for connection...")
                    waitForSocketConnection(socket, callback);
                }
            },
        10);
    }

    function sendRequest(request, withCallback) {
        if (withCallback) {
            var defer = $q.defer();
            var callbackId = getCallbackId();
            callbacks[callbackId] = {
                time: new Date(),
                cb:defer
            };

            request.callback_id = callbackId;
        }

        console.log('Sending request', request);

        waitForSocketConnection(ws, function(){
            console.log("message sent!!!");
            ws.send(JSON.stringify(request));
        });

        if (withCallback) {
            return defer.promise;
        }
    }

    function listener(data) {
        var messageObj = data;
        console.log("Received data from websocket:", messageObj);

        if(callbacks.hasOwnProperty(messageObj.callback_id) && messageObj.type != "msg") {
            console.log("Received callback nr:",messageObj.callback_id);
            callbacks[messageObj.callback_id].cb.resolve(messageObj);
            delete callbacks[messageObj.callback_id];
        }
        else {
            $rootScope.$broadcast('msg', messageObj)
        }
    }

    function getCallbackId() {
        currentCallbackId += 1;
        if(currentCallbackId > 10000) {
            currentCallbackId = 0;
        }
        return currentCallbackId;
    }

    Service.register = function(username) {
        var request = {
            name : username,
            type : 'reg'
        }

        return sendRequest(request, true);
    }

    Service.getUserKey = function(targetuser) {
        if (targetuser) {
            var request = {
                target : targetuser,
                type : 'con'
            }
        }

        return sendRequest(request, true);
    }

    Service.sendMessage = function(message) {
        var request = {
            text : message,
            type : 'msg'
        }

        sendRequest(request, false);
    }
    return Service;
}])

function ChatCtrl($scope, $MyService) {
    $scope.messages = [];
    $scope.name = "";
    $scope.target = "";
    $scope.registered = false;
    $scope.connected = false;

    $scope.$on('msg', function(event, args) {
        if (args.type === "user disconnected") {
            console.log("user is off, target disabled");
            $scope.connected = false;
            $scope.targetuser = "";
        }
    
        console.log("new message event", args);
        $scope.messages.push(args);
        $scope.$digest();
    });
    
    $scope.$on('close', function() {
        console.log("ws closed");
        $scope.messages.push({name: "browser", text: "lost connection to local cryptochat programm, restart it!"});
        $scope.registered = false;
        $scope.connected = false;
        $scope.$digest();
    });

    $scope.register = function() {
        if ($scope.username) {
            $scope.registered = true;
            var c = $MyService.register($scope.username);

            c.then(function(data) {
                console.log("register response: ", data);
                if (data.type === "success") {
                    console.log("register response: ", data);
                }
                else {
                    $scope.registered = false;
                }
                $scope.messages.push(data);
            }, function(data) {
                console.log("then error: ", data);
                $scope.registered = false;
            }, function(data) {
                console.log("then note ", data);
                $scope.registered = false;
            });
        }
    }

    $scope.requestKey = function() {
        if ($scope.targetuser) {
            $scope.connected = true;
            var c = $MyService.getUserKey($scope.targetuser);

            c.then(function(data) {
                if (data.type === "success") {
                    console.log("request key response: ", data);
                }
                else {
                    $scope.connected = false;
                }
                $scope.messages.push(data);
            }, function(data) {
                console.log("then error: ", data);
                $scope.connected = false;
            }, function(data) {
                console.log("then note ", data);
                $scope.connected = false;
            });
        }
    }

    $scope.sendmessage = function() {
        if ($scope.usermsg) {
            $MyService.sendMessage($scope.usermsg);
            $scope.messages.push({name: $scope.username, text: $scope.usermsg});
            $scope.usermsg = "";
        }
    }
}

chatApp.controller('ChatCtrl', ['$scope', 'MyService', ChatCtrl]);