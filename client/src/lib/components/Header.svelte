<header>
    <div>
        <h1>FORUM NAME!</h1>
        <nav>
            <ul>
                <li>
                    <a href="/login">Login</a>
                </li>
                <li>
                    <a href="/logout">Register</a>
                </li>
            </ul>
        </nav>
    </div>
</header>

<style>
    header {
        width: 100%;
        height: 48px;
        position: sticky;
        top: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: var(--bg-secondary);
    }

    header h1 {
        margin: 0;
    }

    div {
        width: var(--app-width);
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
    }

    nav {
        display: flex;
        justify-content: flex-end;
        align-items: center;
    }

    nav > ul {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
        padding: 0;
        list-style-type: none;
    }

    ul > li {
        margin: 0 8px 0 8px;
    }
</style>