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


function signOut() {
    document.cookie = "ssid=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    window.location.reload();
}

function no_implement_sign() {
    alert('Temporary only admins can do this');
}


function getData(form: HTMLFormElement) {
    const data = new FormData(form);

    return data;
}

async function requestSign(path: string, data: object) {
    const resp = fetch('/' + path, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify(data)
    });

    return resp;
}

document.getElementById("signInForm")!.addEventListener("submit", function (e) {
    e.preventDefault();

    let data: { [key: string]: any } = {};

    getData(e.target as HTMLFormElement).forEach(
        (value, key) => data[key] = value
    );

    requestSign('sign-in', data).then((resp) => {
        switch (resp.status) {
            case 401:
                alert('Wrong login or password');
                break;
            default:
                window.location.reload();
        }
    });

});

document.getElementById("signUpForm")!.addEventListener("submit", function (e) {
    e.preventDefault();

    let data: { [key: string]: any } = {};

    getData(e.target as HTMLFormElement).forEach(
        (value, key) => data[key] = value
    );

    requestSign('sign-up', data).then((resp) => {
        switch (resp.status) {
            case 409:
                alert('Failed to create user');
                break;
            case 400:
                alert('Failed to create user');
                break;
            default:
                requestSign('sign-in', data).then((resp) => {
                    switch (resp.status) {
                        case 401:
                            alert('Wrong login or password');
                            break;
                        default:
                            window.location.reload();
                    }
                });
        }
    });
});
