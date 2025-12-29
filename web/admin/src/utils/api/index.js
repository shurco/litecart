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
    
    if (response.status === 204) {
      return { success: true };
    }
    
    const text = await response.text();
    const data = text ? JSON.parse(text) : {};
    
    if (!response.ok) {
      return {
        success: false,
        message: data.message || "Request failed",
        result: data.result || data,
      };
    }
    
    return data;
  } catch (error) {
    if (error instanceof SyntaxError) {
      return {
        success: false,
        message: "Invalid response format",
        result: error.message,
      };
    }
    return {
      success: false,
      message: "Network error",
      result: error.message,
    };
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
