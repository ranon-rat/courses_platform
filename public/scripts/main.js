"use strict";
const url = new URL(window.location.href);
const opt = url.searchParams.get('id');
if (url.pathname === '/' && opt === null) {
    window.location.href = '?id=1';
}
async function search(query) {
    const resp = await fetch('/api?page=1', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        }
    });
    const data = await resp.json();
    const keys = query.split(' ');
    return (data.publications || []).filter((item) => keys.every((key) => item.title.toLowerCase().includes(key)));
}
const input = document.getElementById('search');
input.addEventListener('input', function () {
    const group = document.getElementById('search-group');
    group.innerHTML = '';
    // @ts-ignore
    const query = this.value.toLowerCase();
    if (query.length < 2)
        return group.innerHTML = '';
    search(query).then((data) => {
        for (let i = 0; i < data.length && i < 5; i++) {
            const item = data[i];
            group.innerHTML += `
                <li class="form-control p-0 m-1">
                    <a class="nav-link active" aria-current="page" href="/?id=${item.id}">${item.title}</a>
                </li>
            `;
        }
    });
});
input.addEventListener('blur', function () {
    const group = document.getElementById('search-group');
    setTimeout(() => group.innerHTML = '', 200);
});
// tag, title
const query = document.querySelectorAll('h1, h2, h3, h4, h5, h6');
const sidebar = document.getElementById('sidebar');
for (const item of query) {
    const title = item.textContent || '';
    item.setAttribute('id', title);
    sidebar.innerHTML += `
        <li class="list-group-item">
            <a class="list-group-item list-group-item-action" href="#${title}">${title}</a>
        </li>
    `;
}
const signout = document.getElementById('signout');
signout?.addEventListener('click', function () {
    document.cookie = "ssid=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
});
function no_implement_sign() {
    alert('Temporary only admins can do this');
}
