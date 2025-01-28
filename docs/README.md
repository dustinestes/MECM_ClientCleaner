

TODO: Build Documentation Template

```
project-root/
  ├── api/                    # API-related code (e.g., REST or gRPC)
  │   ├── handler/
  │   │   ├── handler.go      # HTTP request handlers
  │   │   └── ...
  │   ├── middleware/
  │   │   ├── middleware.go   # Middleware for HTTP requests
  │   │   └── ...
  │   └── ...
  ├── cmd/
  │   ├── your-app-name/
  │   │   ├── main.go         # Application entry point
  │   │   └── ...             # Other application-specific files
  │   └── another-app/
  │       ├── main.go         # Another application entry point
  │       └── ...
  ├── configs/                # Configuration files for different environments
  │   ├── development.yaml
  │   ├── production.yaml
  │   └── ...
  ├── deploy                  # IaaS, PaaS, system, and container orchestration deployment configurations and templates
  ├── docs/                   # Project documentation
  ├── github/                 # Files used for the github aspect of the project
  |   ├── assets/             # Logo, images, screenshots, etc. for github readme
  |   ├── githooks/           # Git hooks
  |   ├── website/            # Files for the project's website (if not using GitHub pages)
  ├── internal/               # Private application and package code
  │   ├── config/
  │   │   ├── config.go       # Configuration logic
  │   │   └── ...
  │   ├── database/
  │   │   ├── database.go     # Database setup and access
  │   │   └── ...
  │   └── ...
  ├── pkg/                    # Public, reusable packages
  │   ├── mypackage/
  │   │   ├── mypackage.go    # Public package code
  │   │   └── ...
  │   └── ...
  ├── scripts/                # Build, deployment, and maintenance scripts
  │   ├── build.sh
  │   ├── deploy.sh
  │   └── ...
  ├── test/                   # Unit and integration tests
  │   ├── unit/
  │   │   ├── ...
  │   └── integration/
  │       ├── ...
  ├── web/                    # Front-end web application assets
  │   ├── static/
  │   │   ├── css/
  │   │   ├── js/
  │   │   └── ...
  │   └── templates/
  │       ├── index.html
  │       └── ...
  ├── .gitignore              # Gitignore file
  ├── go.mod                  # Go module file
  ├── go.sum                  # Go module dependencies file
  └── README.md               # Project README
```


- cmd/: This directory contains application-specific entry points (usually one per application or service). It's where you start your application.
- internal/: This directory holds private application and package code. Code in this directory is not meant to be used by other projects. It's a way to enforce access control within your project.
- pkg/: This directory contains public, reusable packages that can be used by other projects. Code in this directory is meant to be imported by external projects.
- api/: This directory typically holds HTTP or RPC API-related code, including request handlers and middleware.
- web/: If your project includes a front-end web application, this is where you'd put your assets (CSS, JavaScript, templates, etc.).
- scripts/: Contains scripts for building, deploying, or maintaining the project.
- configs/: Configuration files for different environments (e.g., development, production) reside here.
- tests/: Holds unit and integration tests for your code.
- docs/: Project documentation, such as design documents or API documentation.