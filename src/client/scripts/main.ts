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

input.addEventListener('blur', function () {
    const group = document.getElementById('search-group')!;
    setTimeout(() => group.innerHTML = '', 200);
})

// tag, title
let titles: [string, string][] = [];

for (const tag of ['h1', 'h2', 'h3', 'h4', 'h5', 'h6']) {
    const elements = document.getElementsByTagName(tag) as HTMLCollectionOf<HTMLElement>;

    for (let i = 0; i < elements.length; i++) {
        const element = elements[i]; // @ts-ignore
        element.setAttribute('id', `${element.innerHTML}`);
        titles.push([element.id, element.innerText]);
    }
}

const sidebar = document.getElementById('sidebar')!;

for (const [id, title] of titles) {
    sidebar.innerHTML += `
        <li class="list-group-item">
            <a class="list-group-item list-group-item-action" href="#${id}">${title}</a>
        </li>
    `;
}
