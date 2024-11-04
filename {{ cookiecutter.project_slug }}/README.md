# Candorship

A radical transparency platform helping organizations build trust through goal
tracking, decision documentation, and progress visualization.

## About

Candorship helps teams and organizations implement radical transparency by providing:
- Clear OKR tracking and visualization
- Structured decision documentation
- Real-time progress monitoring
- Cross-team alignment tools

## Core Features

### Goal Transparency Hub
- Comprehensive OKR management system
- Goal hierarchy visualization
- Automated check-ins and updates
- Team progress tracking
- Custom reporting
- Dependencies mapping

### Decision Dashboard
- Standardized decision documentation
- Clear ownership and stakeholder mapping
- Timeline visualization
- Impact assessment
- Feedback collection
- Performance analytics

## Getting Started

Candorship is built using [Echo][1] and [TailwindCSS][2].

### Running the Development Server

1. Ensure that you have Go and TailwindCSS installed. We recommend using the
[pre-built binaries][3] for TailwindCSS instead of installing it via npm.
2. Run `go get .` to install Go dependencies
3. In a separate terminal, run `tailwindcss -i static/css/input.css -o static/css/app.css --watch --minify` to build and minify the CSS files
4. Run `air` to start the development server with autoreload enabled

### License

Copyright 2024 Perch Labs Pte Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


[1]: https://echo.labstack.com/
[2]: https://tailwindcss.com
[3]: https://github.com/tailwindlabs/tailwindcss/releases