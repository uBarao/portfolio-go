const containers = document.querySelectorAll('.project-card');

containers.forEach(container => {
  const card = container.querySelector('.project-hover-card');

  container.addEventListener('mousemove', (event) => {
    const cardWidth = card.offsetWidth;
    const cardHeight = card.offsetHeight;
    const mouseX = event.clientX;
    const mouseY = event.clientY;

    const viewportWidth = window.innerWidth;
    const viewportHeight = window.innerHeight;

    let left = mouseX;
    let top = mouseY;

    if (mouseX + cardWidth > viewportWidth) {
      left = viewportWidth - cardWidth;
    }

    if (mouseY + cardHeight > viewportHeight) {
      top = viewportHeight - cardHeight;
    }

    card.style.left = `${left}px`;
    card.style.top = `${top}px`;
    card.style.visibility = 'visible';
    card.style.opacity = '1';
  });

  container.addEventListener('mouseleave', () => {
    card.style.visibility = 'hidden';
    card.style.opacity = '0';
  });
});
