package cloudelements

var jsonMetadata = `{
  "name": "cloudelements-instance",
  "version": "0.0.1",
  "title": "Invoke Cloud Elements Instance",
  "description": "Simple Cloud Elements Activity",
  "inputs":[
    {
      "name": "organizationKey",
      "type": "string",
      "required": true
    },
    {
      "name": "userKey",
      "type": "string",
      "required": "true"
    },
    {
      "name": "elementKey",
      "type": "string",
      "required": false
    },
    {
      "name": "method",
      "type": "string",
      "required": true
    },
    {
      "name": "uri",
      "type": "string",
      "required": true
    },
    {
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "queryParams",
      "type": "params"
    },
    {
      "name": "content",
      "type": "any"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "any"
    }
  ]
}
`
