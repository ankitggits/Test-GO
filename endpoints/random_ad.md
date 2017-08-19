# Advertisement Resources

    GET RandomAd

## Description
Returns randomly picked Ad from any category.

***

## Example
* **URL**

   /service

* **Method:**

  `GET`
  
*  **URL Params**

   None

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />

    

* **Return** 
    ``` json
    {
        "response":{
            "ad_key":"Ubisoft_Inc_1",
            "ad_provider":"Ubisoft_Inc.",
            "ad_text":"http://ubisoft.com"
        },
        "requestPath":"/service",
        "executionTime":"0s"
    }