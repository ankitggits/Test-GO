# Advertisement Resources

    GET SearchTextCategoryAd

## Description
Api will find a random ad from given category and search text in URL path param , in case of unavailability return 404 http status code

##### Note: search text can contain adKey or adProvider. adKey has given priority over provider

***

## Example
* **Request**

    http://localhost:8090/service/{:category}/{:keyOrProvider}
    

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
      `category=[string]`
      
      `keyOrProvider=[string]`

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    
* **Return** 
``` json
{
    "response": {
      "ad_key": "TFG_AS_2",
        "ad_provider": "The_Future_Group_AS",
        "ad_text": "https://www.futureuniverse.com"
    },
    "requestPath": "/service/IMR",
    "executionTime": "0s"
}
