# init

1. `make init`

1. setting

    You have to set your nature remo access token and your light signal id in `.env` file.
    You can check your light signal id to exec `make get_appliances`.

1. make Docker image

    '''
    docker build -t nature_remo_auto_change_light .
    '''

1. check

    ```
    docker run --rm nature_remo_auto_change_light
    ```
