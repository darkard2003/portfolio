document.addEventListener('alpine:init', () => {
    // Dark Mode Store
    Alpine.store('darkMode', {
        on: localStorage.getItem('darkMode') === null ? true : localStorage.getItem('darkMode') === 'true',
        toggle() {
            this.on = !this.on;
            localStorage.setItem('darkMode', this.on);

            // Force class update for immediate feedback
            if (this.on) {
                document.documentElement.classList.remove('light');
            } else {
                document.documentElement.classList.add('light');
            }
        }
    });

    // Mobile Menu Store
    Alpine.store('mobileMenu', {
        isOpen: false,
        toggle() { this.isOpen = !this.isOpen },
        close() { this.isOpen = false }
    });

    // Project Filter Store
    Alpine.store('projectFilter', {
        selected: 'none',
        toggle(tech) {
            this.selected = this.selected === tech ? 'none' : tech;
        },
        isSelected(tech) {
            return this.selected === tech;
        },
        isVisible(projectTechs) {
            return this.selected === 'none' || projectTechs.includes(this.selected);
        }
    });

    // Navigation Store for Logo/Back Button
    Alpine.store('navigation', {
        path: window.location.pathname,
        isHome() {
            return this.path === '/' || this.path === '';
        },
        updatePath() {
            this.path = window.location.pathname;
        }
    });
});

// HTMX Events for Robustness
document.addEventListener('htmx:afterSwap', (event) => {
    // Update navigation state after HTMX swap
    if (window.Alpine) {
        const nav = Alpine.store('navigation');
        if (nav) nav.updatePath();
    }
});

// Handle History Navigation (Back/Forward)
window.addEventListener('popstate', () => {
    if (window.Alpine) {
        const nav = Alpine.store('navigation');
        if (nav) nav.updatePath();
    }
});
