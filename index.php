<?php
	//default redis port changed in redis.conf
	require "predis/autoload.php"
	PredisAutoloader::register();

	try {
		$redis = new PredisClient();
	}
	catch*(Exception $e) {
		die($e->getMessage());
	}
		
	//Receive the RAW post data via the php://input IO stream.
	$content = file_get_contents("php://input");
	//Make sure that it is a POST request.
	if(strcasecmp($_SERVER['REQUEST_METHOD'], 'POST') != 0){
	    throw new Exception('Request method must be POST!');
	}
	 
	//Make sure that the content type of the POST request has been set to application/json
	// $contentType = isset($_SERVER["CONTENT_TYPE"]) ? trim($_SERVER["CONTENT_TYPE"]) : '';
	// if(strcasecmp($contentType, 'application/json') != 0){
	//     throw new Exception('Content type must be: application/json');
	// }
	 
	//Receive the RAW post data.
	$content = trim(file_get_contents("php://input"));
	 
	//Attempt to decode the incoming RAW post data from JSON.
	$decoded = json_decode($content, true);
	 
	//If json_decode failed, the JSON is invalid.
	if(!is_array($decoded)){
	    throw new Exception('Received content contained invalid JSON!');
	}
	 
	//provide that JSON is not null and structured per spec
	if(isset($decoded["endpoint"]) && isset($decoded["data"])) {
		
		//make later code more legible
		$method = $postback["endpoint"]["method"];
		$url = $postback["endpoint"]["url"];
		$data = $decoded["data"];

		//ensure there is data
		if(count($array) > 0) {
			foreach($decoded["data"] as $datum) {
				$postbackObject = (object)array(
					"method"=> $method, 
					"url" => $url,
					"data" => $datum
				);
				$redis->rpush("postback", $postbackObject);
			}
		}
	}
?>

