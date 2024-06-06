# GoOk
Check if value is valid:
- email
    ```
    gook.Email("email@example.com").Ok() // return boolean
    gook.Email("email@example.com").Error() // return error
    ```
- password
    ```
    gook.Password("@BCd3fgh1").Ok() // return boolean
    gook.Password("@BCd3fgh1").Error() // return error
    ```
- phone
    ```
    gook.Phone("0123456789").Ok() // return boolean
    gook.Phone("0123456789").Error() // return error
    ```

