document.addEventListener('alpine:init', () => {
    Alpine.store('darkMode', {
        on: localStorage.getItem('darkMode') === null ? true : localStorage.getItem('darkMode') === 'true',
        toggle() {
            const switchTheme = () => {
                this.on = !this.on;
                localStorage.setItem('darkMode', this.on);

                if (this.on) {
                    document.documentElement.classList.remove('light');
                } else {
                    document.documentElement.classList.add('light');
                }
            };

            if (!document.startViewTransition) {
                switchTheme();
                return;
            }

            document.startViewTransition(switchTheme);
        }
    });

    Alpine.store('mobileMenu', {
        isOpen: false,
        toggle() { this.isOpen = !this.isOpen },
        close() { this.isOpen = false }
    });

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

document.addEventListener('htmx:afterSwap', (event) => {
    if (window.Alpine) {
        const nav = Alpine.store('navigation');
        if (nav) nav.updatePath();
    }
});

window.addEventListener('popstate', () => {
    if (window.Alpine) {
        const nav = Alpine.store('navigation');
        if (nav) nav.updatePath();
    }
});
