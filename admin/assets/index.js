document.addEventListener('DOMContentLoaded', function() {
    const toggleButton = document.querySelector('.toggle-sidebar');
    const sidebar = document.querySelector('.sidebar');
    const content = document.querySelector('.content');

    toggleButton.addEventListener('click', function() {
        sidebar.classList.toggle('collapsed');
        content.classList.toggle('expanded');
    });
});
