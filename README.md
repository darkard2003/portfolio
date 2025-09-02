# Portfolio Website

**Hosted Link:** [kaushikchowdhury.dev](https://kaushikchowdhury.dev)

This is the source code for my personal portfolio website, built to showcase my projects, skills, and experience as a developer.

## About the Project

- **Blazing fast**: The site is optimized for quick loading and smooth navigation, ensuring a seamless experience for visitors.
- **Responsive design**: Built with TailwindCSS, the layout adapts beautifully to all screen sizes, from mobile to desktop.
- **Type-safe templates**: Uses Templ for Go, providing robust, type-safe HTML generation and component composition.
- **Single source of truth**: All content is managed via a single `data.json` file, making updates simple and consistent.
- **Modern stack**: Powered by Go, Templ, TailwindCSS, and Alpine.js for lightweight interactivity.

This portfolio is a reflection of my skills and design philosophy—minimal, efficient, and focused on delivering a great user experience.

## File Structure

```
.
├── .github/             # GitHub related configurations
├── cmd/                 # Main application entry points
│   └── web/             # Web application server
├── internals/           # Internal packages and logic
│   ├── handelers/       # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models and structures
│   └── utils/           # Utility functions
├── static/              # Static assets (CSS, JS, images)
│   ├── css/             # Compiled CSS files
│   └── scripts/         # JavaScript files
├── web/                 # Web-related source files
│   ├── css/             # Tailwind CSS source
│   └── view/            # Templ HTML templates
│       ├── components/  # Reusable UI components
│       ├── layout/      # Base HTML layouts
│       └── pages/       # Page-specific templates
├── data.json            # Application data
├── Dockerfile           # Docker container definition
├── go.mod               # Go module dependencies
├── LICENSE              # Project license
├── makefile             # Build automation scripts
├── package.json         # Node.js package definitions
├── README.md            # Project README
└── ...                  # Other configuration files (e.g., .air.toml, .dockerignore, .gitignore, fly.toml)
```

---

© 2025 darkard2003. All rights reserved.
