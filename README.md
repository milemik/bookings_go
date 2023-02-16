# Bookings web app in GoLang

Just a test project.

- Built in GoVersion - 1.92
- Uses [chi router](https://github.com/go-chi/chi/v5)
- Usess [SCS session middleware](https://github.com/alexedwards/scs/v2)
- Usess [nosurf](https://github.com/justinas/nosurf)


This is just project for practice GoLang

## SETUP FOR DEVELOPMENT

1. Make sure you have docker installed
2. Make sure you have soda installed [soda/pop](https://gobuffalo.io/documentation/database/soda/)
    - Add soda to path by adding this to your .zshrc or .bashrc or what bash you are using (ON LINUX)
    ```shell
    export PATH="$HOME/go/bin:$PATH"
    ```
2. Run
    ```shell
    # This will start postgres container
    bash run_db.sh
    ```
3. Run migrations
    ```shell
    soda migrate
    ```
###NOTE:make sure to run postgresql before runing migrations