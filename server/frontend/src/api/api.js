    const API_URL = "http://localhost:8000";

    export async function GetArticles() {
        const response = await fetch(`${API_URL}/article`);
        return await response.json();
    }

    export async function CreateArticles(article) {
        const response = await fetch(`${API_URL}/admine`, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(article)
        });
        return await response.json();
    }