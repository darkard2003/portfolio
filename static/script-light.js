// Lightweight icon replacement for Feather Icons
// Load SVG sprite and replace data-feather attributes

document.addEventListener('DOMContentLoaded', function() {
  // Load and inject SVG sprite
  fetch('/static/icons.svg')
    .then(response => response.text())
    .then(svgSprite => {
      // Inject SVG sprite into page
      const div = document.createElement('div');
      div.innerHTML = svgSprite;
      document.body.insertBefore(div, document.body.firstChild);
      
      // Replace all data-feather elements
      replaceIcons();
    });
});

function replaceIcons() {
  const elements = document.querySelectorAll('[data-feather]');
  
  elements.forEach(element => {
    const iconName = element.getAttribute('data-feather');
    const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    const use = document.createElementNS('http://www.w3.org/2000/svg', 'use');
    
    // Set SVG attributes
    svg.setAttribute('width', '24');
    svg.setAttribute('height', '24');
    svg.setAttribute('viewBox', '0 0 24 24');
    svg.setAttribute('fill', 'none');
    svg.setAttribute('stroke', 'currentColor');
    svg.setAttribute('stroke-width', '2');
    svg.setAttribute('stroke-linecap', 'round');
    svg.setAttribute('stroke-linejoin', 'round');
    
    // Link to symbol
    use.setAttributeNS('http://www.w3.org/1999/xlink', 'href', `#${iconName}`);
    
    svg.appendChild(use);
    element.parentNode.replaceChild(svg, element);
  });
}

// Mobile navigation toggle
function toggleMobileNav() {
  const navMenu = document.getElementById('nav-menu');
  navMenu.classList.toggle('mobile-hidden');
}

// Close mobile nav when clicking on a link
document.addEventListener('DOMContentLoaded', function() {
  document.querySelectorAll('.nav-menu a').forEach(link => {
    link.addEventListener('click', () => {
      document.getElementById('nav-menu').classList.remove('mobile-hidden');
    });
  });

  // Close mobile nav when clicking outside
  document.addEventListener('click', (e) => {
    const nav = document.querySelector('nav');
    const navMenu = document.getElementById('nav-menu');
    if (!nav.contains(e.target)) {
      navMenu.classList.remove('mobile-hidden');
    }
  });
});
