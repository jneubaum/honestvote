go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "3059301306072a8648ce3d020106082a8648ce3d03010703420004d5f116c3064bf914b02ed5b6a888607dde5437dbebe99fd9a518440e9fd730fa2ab6edb15fafa2ec68e3a6d1387450966582479064ed81e6f66ce4f08abc8d92" --public-key "307702010104204266aa2680e36f7bd93c33d69f29bf93b52f76208a2a657a77b68d533542d38aa00a06082a8648ce3d030107a14403420004d5f116c3064bf914b02ed5b6a888607dde5437dbebe99fd9a518440e9fd730fa2ab6edb15fafa2ec68e3a6d1387450966582479064ed81e6f66ce4f08abc8d92" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "3059301306072a8648ce3d020106082a8648ce3d030107034200045140ef4bf40f539d3015ec4b2ff28f0926aa57ec95653c161ae4348abd40a617d729c5bcffa28c48d63124b0c371b84626fa62b8fce7a2fc1d6c9f21b9ab5dc5" --public-key "3077020101042003eebc983b01e58715d4dfde333216245b3f5c674db74c333b33d1476b17f9fca00a06082a8648ce3d030107a144034200045140ef4bf40f539d3015ec4b2ff28f0926aa57ec95653c161ae4348abd40a617d729c5bcffa28c48d63124b0c371b84626fa62b8fce7a2fc1d6c9f21b9ab5dc5" & \