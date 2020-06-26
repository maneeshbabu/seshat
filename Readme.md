# Seshat

She knows everything

## API Documentation

- create agent
    - [post] /v1/api/agent
        - request
            ```json
            {
                "name": "example_mrss"
            }
            ```
        - response
            - 201
            ```json 
            {"code": "agent_code"}
            ```
            - 400
            ```json 
            {"error": "bad request"}
            ```
            - 403
            ```json 
            {"error": "message"}
            ```

- create job
    - [post] /v1/api/agent/:code/job
        - request
            ```json
            {
                "blip_domain": "example.amagi.tv",
                "platform": "example",
                "asset_id": "asset_id",
                "channel_name": "channel_name",
                "type": "asset_ingest",
                "additional_info": {},
            }
            ```
        - response
            - 201
            ```json 
            {"id": "status_id"}
            ```
            - 400
            ```json 
            {"error": "bad request"}
            ```
            - 403
            ```json 
            {"error": "message"}
            ```
- create status for the job
    - [post] /v1/api/agent/:code/job/:id/status/:state
        - request
            ```json
            {
                "status": "copying",
                "percentage": "50.65"
            }
            ```
        - response
            - 201
            ```json 
            {"id": "status_id"}
            ```
            - 400
            ```json 
            {"error": "bad request"}
            ```
            - 403
            ```json 
            {"error": "message"}
            ```

- list agents
    - [get] /v1/api/agent
        - response
            - 200
            ```json 
            [
                {
                    "code": "agent_code", 
                    "name": "agent_name" 
                },
            ]
            ```
- list jobs
    - [get] /v1/api/agent/:code/job
        - response
            - 200
            ```json 
            [
                {
                    "id": "job_id", 
                    "blip_domain": "blip_domain" ,
                    "platform": "platform", 
                    "asset_id": "asset_id",
                    "type": "asset_ingest", 
                    "additional_info": "additional_info"
                },
            ]
            ```
            
            - 404
            ```json
            {"error": "NOT FOUND"}
            ```
- list job status
    - [get] /v1/api/agent/:code/job/:id/status
        - 200
            ```json
            {
                "id": "job_id", 
                "blip_domain": "blip_domain" ,
                "platform": "platform", 
                "asset_id": "asset_id",
                "type": "asset_ingest", 
                "additional_info": "additional_info",
                "status": [
                    {
                        "id": "status_id",
                        "status":"status",
                        "percentage": "",
                    }
                ]
            }
            ```
        - 404
            ```json
            {"error": "NOT FOUND"}
            ```

## Nomenclature
- AGENT - agent is the one perform jobs
    ```
    code - agent code
    name - agent name
    ```
- JOB - job is action performing by agent
    ```
    - job_id - unique identifier generator by hermes
    - blip_domain - domain for which the job is performed for
    - platform - platform for which action is performed for, normally this is same as blip. But the assets are hosted by a different vendor platform name must be that vendor
    - channel_name - channel name
    - type - job type - (asset_ingest, meta_ingest, subtitle_ingest...)
    - asset_id - asset id
    - additional_info - additional information
    - status - status of the job 
    ```
- STATUS - status of the job at any point of time
    ```
    - job_id - identifier to the job
    - status - status of the job
    - percentage - percentage of completion 
    ```


## Queries
1. How can we trigger failed jobs from seshat?

    seshat will be uploading job json to the agents trigger location. This will be causing a re trigger of job based on cron frequency.