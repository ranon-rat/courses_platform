const query = document.querySelectorAll('h1, h2, h3, h4, h5, h6');
const sidebar = document.getElementById('sidebar');

if (sidebar) for (const item of query) {
    const title = item.textContent || '';

    if (item.classList.contains('modal-title')) continue;
    item.setAttribute('id', title);

    sidebar.innerHTML += `
        <li class="list-group-item">
            <a class="list-group-item list-group-item-action" href="#${title}">${title}</a>
        </li>
    `;
}
