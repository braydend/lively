import React, {useState} from 'react';
import styles from './App.module.css';
import { Header, Message } from "./components";

function App() {
    const [result, setResult] = useState();

    window.backend.getMessage().then((result) => setResult(result));

    return (
        <main className={styles.App}>
            <Header />
            <Message message={result} />
        </main>
    );
}

export default App;
