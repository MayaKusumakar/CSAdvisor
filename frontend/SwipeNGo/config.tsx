

const apiConfig = {
    "endpoint": "http://localhost:8080",
    "available_apis": {
        "getEventsForUser": {
            "url": "http://localhost:8080",
            "method": "POST",
            "headers": {},
            "body": {
                "user_id": "String",
                "user_location": {
                    "latitude": "String",
                    "longitude": "String"
                }
            }
        }
    }
}