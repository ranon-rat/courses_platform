const url = new URL(window.location.href);
const opt = url.searchParams.get('id');


if (url.pathname === '/' && opt === null) {
    window.location.href = '?id=1';
}


interface Response {
    publications: {
        title: string;
        id: string;
    }[];
}

async function search(query: string): Promise<Response['publications']> {
    const resp = await fetch('/api?page=1', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        }
    });

    const data: Response = await resp.json();
    const keys = query.split(' ');

    return (data.publications || []).filter((item) =>
        keys.every((key) => item.title.toLowerCase().includes(key))
    )
}
const input = document.getElementById('search')!

input.addEventListener('input', function () {
    const group = document.getElementById('search-group')!;
    group.innerHTML = '';

    // @ts-ignore
    const query = this.value.toLowerCase();

    if (query.length < 2) return group.innerHTML = '';

    search(query).then((data) => {
        for (let i = 0; i < data.length && i < 5; i++) {
            const item = data[i];

            group.innerHTML += `
                <li class="form-control p-0 m-1">
                    <a class="nav-link active" aria-current="page" href="/?id=${item.id}">${item.title}</a>
                </li>
            `;
        }
    })
})

function signOut() {
    document.cookie = "ssid=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    window.location.reload();
}
