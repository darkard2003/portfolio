document.addEventListener('alpine:init', () => {
    // Dark Mode Store
    Alpine.store('darkMode', {
        on: localStorage.getItem('darkMode') === null ? true : localStorage.getItem('darkMode') === 'true',
        toggle() {
            this.on = !this.on;
            localStorage.setItem('darkMode', this.on);
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
});
