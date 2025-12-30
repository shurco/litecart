# Design Customization and Deployment

## Changing the Site Design

### Frontend Structure

The site frontend is located in the `web/site/` directory and uses:
- **SvelteKit** - framework for creating SPA
- **TailwindCSS v4** - for styling
- **TypeScript** - for type safety

### Main Files for Customization

#### 1. Global Styles

The `web/site/src/app.css` file contains global styles and TailwindCSS configuration:

```css
@import 'tailwindcss';
@plugin "@tailwindcss/forms";

/* Additional styles for product description */
.prod_desc > * + * {
  margin-top: 0.75em;
}
```

Here you can:
- Change the color scheme via TailwindCSS
- Add custom CSS classes
- Configure typography

#### 2. Components

Main components are located in `web/site/src/lib/components/`:

- **Header.svelte** - site header with logo, social networks, and cart
- **Footer.svelte** - site footer with page links
- **ProductCard.svelte** - product card
- **FormButton.svelte** - form button

Example of modifying Header:

```svelte
<!-- web/site/src/lib/components/Header.svelte -->
<header class="bg-white">
  <!-- Change bg-white to desired color, e.g., bg-blue-50 -->
  <div class="mx-auto flex h-16 max-w-screen-xl items-center gap-8 px-4 sm:px-6 lg:px-8">
    <!-- Your custom content -->
  </div>
</header>
```

#### 3. Layout

The main layout is located in `web/site/src/lib/layouts/MainLayout.svelte`. Here you can change the overall page structure.

#### 4. Pages

Pages are located in `web/site/src/routes/`:
- `+page.svelte` - home page
- `[slug]/+page.svelte` - content pages
- `products/[slug]/+page.svelte` - product page
- `cart/+page.svelte` - cart

### Design Modification Process

1. **Development Mode**

   Start the development server to view changes in real-time:

   ```bash
   # Start litecart server (in one terminal)
   ./litecart serve

   # Start frontend dev server (in another terminal)
   cd web/site
   bun run dev
   ```

   The site will be available at `http://localhost:5273`

2. **Making Changes**

   - Modify components in `web/site/src/lib/components/`
   - Modify styles in `web/site/src/app.css`
   - Modify pages in `web/site/src/routes/`

3. **Building for Production**

   After making changes, build the frontend:

   ```bash
   cd web/site
   bun run build
   ```

   The built files will be in `web/site/build/` and will be automatically embedded into the binary on the next litecart build.

4. **Rebuilding litecart**

   If you're using a compiled binary, rebuild it:

   ```bash
   go build -o litecart ./cmd/main.go
   ```

### Customization via TailwindCSS

You can configure colors, fonts, and other parameters via TailwindCSS. Create a `web/site/tailwind.config.js` file if it doesn't exist:

```javascript
/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#eff6ff',
          500: '#3b82f6',
          700: '#1d4ed8',
        },
      },
    },
  },
  plugins: [],
}
```

Then use these colors in components:

```svelte
<div class="bg-primary-500 text-white">
  Custom color
</div>
```

## Deployment on a Separate Server with Nginx

### Option 1: Frontend and API on the Same Domain

This is the standard configuration where the frontend and API work on the same domain through Nginx.

#### Step 1: Build Frontend

```bash
cd web/site
bun install
bun run build
```

#### Step 2: Configure Nginx

Create Nginx configuration `/etc/nginx/sites-available/litecart`:

```nginx
server {
    listen 80;
    server_name yourdomain.com;
    
    # Maximum upload file size
    client_max_body_size 20M;
    
    # Gzip compression
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css text/xml text/javascript application/x-javascript application/xml+rss application/json;

    # Frontend static files
    root /path/to/litecart/web/site/build;
    index index.html;

    # API and uploads are proxied to litecart server
    location /api {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }

    location /uploads {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # Admin panel
    location /_/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_cache_bypass $http_upgrade;
    }

    # SPA routing - all other requests to index.html
    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

#### Step 3: Activate Configuration

```bash
sudo ln -s /etc/nginx/sites-available/litecart /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

#### Step 4: Start litecart Server

```bash
./litecart serve --http 127.0.0.1:8080 --no-site
```

The `--no-site` flag disables the embedded frontend since we're using a separate one.

### Option 2: Frontend and API on Different Domains/Servers

If the frontend should work on a separate domain or server, you need to configure CORS and specify the API address.

#### Step 1: Configure API Server on a Different Address

Create a configuration file for API URL. In `web/site/src/lib/utils/api.ts` you can add support for environment variables:

```typescript
// web/site/src/lib/utils/api.ts

// Base API URL (can be set via environment variable)
const API_BASE_URL = import.meta.env.VITE_API_URL || '';

async function handleRequest<T = any>(url: string, options: RequestOptions): Promise<ApiResponse<T>> {
  try {
    // Add base URL if specified
    const fullUrl = API_BASE_URL ? `${API_BASE_URL}${url}` : url;
    const response = await fetch(fullUrl, options)
    // ... rest of the code
  }
}
```

#### Step 2: Configure Environment Variable

Create a `.env` file in `web/site/`:

```env
VITE_API_URL=https://api.yourdomain.com
```

Or during build:

```bash
VITE_API_URL=https://api.yourdomain.com bun run build
```

#### Step 3: Configure CORS on API Server

Make sure the litecart server allows requests from your frontend domain. This is usually configured in the server code (check `internal/app.go`).

#### Step 4: Nginx Configuration for Frontend

```nginx
server {
    listen 80;
    server_name frontend.yourdomain.com;
    
    root /path/to/litecart/web/site/build;
    index index.html;

    # SPA routing
    location / {
        try_files $uri $uri/ /index.html;
    }

    # Static files
    location /assets {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

#### Step 5: Nginx Configuration for API

```nginx
server {
    listen 80;
    server_name api.yourdomain.com;
    
    client_max_body_size 20M;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # CORS headers (if not configured in the application)
        add_header 'Access-Control-Allow-Origin' 'https://frontend.yourdomain.com' always;
        add_header 'Access-Control-Allow-Methods' 'GET, POST, PATCH, DELETE, OPTIONS' always;
        add_header 'Access-Control-Allow-Headers' 'Content-Type, Authorization' always;
        add_header 'Access-Control-Allow-Credentials' 'true' always;
        
        if ($request_method = 'OPTIONS') {
            return 204;
        }
    }
}
```

### HTTPS Configuration

For production, it's recommended to use HTTPS. Use Let's Encrypt:

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com
```

Certbot will automatically update the Nginx configuration to use HTTPS.

### Auto-start litecart

Create a systemd service `/etc/systemd/system/litecart.service`:

```ini
[Unit]
Description=Litecart Server
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/path/to/litecart
ExecStart=/path/to/litecart/litecart serve --http 127.0.0.1:8080 --no-site
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Activate the service:

```bash
sudo systemctl enable litecart
sudo systemctl start litecart
```

## Configuring API Server on a Different Address

### For Development

In the `web/site/vite.config.ts` file, configure the proxy:

```typescript
export default defineConfig({
  plugins: [sveltekit(), tailwindcss()],
  server: {
    port: 5273,
    proxy: {
      '/api': {
        target: 'http://your-api-server.com:8080',
        changeOrigin: true
      },
      '/uploads': {
        target: 'http://your-api-server.com:8080',
        changeOrigin: true
      }
    }
  }
})
```

### For Production

Use the `VITE_API_URL` environment variable as described above in the "Option 2" section.

### Alternative Method via Configuration File

You can create a configuration file `web/site/src/lib/config.ts`:

```typescript
// web/site/src/lib/config.ts

export const config = {
  apiUrl: import.meta.env.VITE_API_URL || (typeof window !== 'undefined' 
    ? window.location.origin 
    : ''),
}

// In api.ts use:
import { config } from '$lib/config'

const fullUrl = config.apiUrl ? `${config.apiUrl}${url}` : url;
```

This allows you to easily change the API address without rebuilding if you load the configuration dynamically.

## Complete Setup Example on a Separate Server

### Directory Structure

```
/var/www/
├── litecart/              # Binary and data
│   ├── litecart          # Executable file
│   ├── litecart.service  # Systemd service
│   ├── lc_base/         # Database
│   ├── lc_uploads/       # Uploaded files
│   └── lc_digitals/     # Digital products
└── litecart-frontend/    # Built frontend
    └── build/           # Static files
```

### Deployment Commands

```bash
# 1. Create directories
sudo mkdir -p /var/www/litecart
sudo mkdir -p /var/www/litecart-frontend

# 2. Copy binary
sudo cp litecart /var/www/litecart/

# 3. Build frontend
cd web/site
bun install
bun run build
sudo cp -r build/* /var/www/litecart-frontend/

# 4. Set permissions
sudo chown -R www-data:www-data /var/www/litecart
sudo chown -R www-data:www-data /var/www/litecart-frontend

# 5. Create systemd service
sudo nano /etc/systemd/system/litecart.service
# (paste configuration from section above)

# 6. Configure Nginx
sudo nano /etc/nginx/sites-available/litecart
# (paste configuration from section above)

# 7. Activate everything
sudo systemctl enable litecart
sudo systemctl start litecart
sudo ln -s /etc/nginx/sites-available/litecart /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

After this, your site will be available at the specified domain.
