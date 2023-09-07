import { start, done } from "nprogress";

export async function apiGet(url) {
  return handleRequest(url, {
    credentials: "include",
    method: "GET",
  });
}

export async function apiPost(url, body) {
  const options = createOptions("POST", body);
  return handleRequest(url, options);
}

export async function apiUpdate(url, body) {
  const options = createOptions("PATCH", body);
  return handleRequest(url, options);
}

export async function apiDelete(url) {
  return handleRequest(url, {
    credentials: "include",
    method: "DELETE",
  });
}

async function handleRequest(url, options) {
  try {
    start();
    const response = await fetch(url, options);
    return response.json();
  } catch (error) {
    console.error(error);
  } finally {
    done();
  }
}

function createOptions(method, body) {
  const options = {
    credentials: "include",
    method,
  };

  if (body) {
    if (Object.keys(body).length > 0) {
      options.body = JSON.stringify(body);
      options.headers = {
        "Content-Type": "application/json",
      };
    } else {
      options.body = body;
    }
  }

  return options;
}
