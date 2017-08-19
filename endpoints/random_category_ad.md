# Advertisement Resources

    GET RandomCategoryAd

## Description
Returns randomly picked Ad by given category.

***

## Example
* **Request**

    http://localhost:8090/service/{:category}
    
* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
      `category=[string]`

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