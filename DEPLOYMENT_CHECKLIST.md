# Deployment Checklist for Docker & CI/CD

## ✅ What's Ready

### 1. Docker Configuration
- ✅ `Dockerfile` - Standard multi-stage build
- ✅ `Dockerfile.frontend-artifact` - Optimized for CI/CD with pre-built frontend
- ✅ `docker-compose.yml` - Ready for local deployment
- ✅ `entrypoint.sh` - Exists and is executable

### 2. GitHub Actions Workflows
- ✅ `.github/workflows/docker.yml` - Docker build and push to Docker Hub & GHCR
- ✅ `.github/workflows/release.yml` - Linux binary builds for multiple architectures
- ✅ `.github/workflows/windows.yml` - Windows binary builds

### 3. Registry Configuration
- ✅ Docker Hub: `zamibd/s-ui` (fixed from `imzami/s-ui`)
- ✅ GitHub Container Registry: `ghcr.io/zamibd/s-ui`
- ✅ Multi-platform builds: amd64, arm64, armv7, armv6, 386

## ⚠️ Required GitHub Secrets

You need to configure the following secrets in your GitHub repository:

### Docker Hub Secrets
1. Go to: Repository Settings → Secrets and variables → Actions → New repository secret
2. Add these secrets:
   - `DOCKER_HUB_USERNAME` - Your Docker Hub username (e.g., `zamibd`)
   - `DOCKER_HUB_TOKEN` - Your Docker Hub access token
     - Get token from: https://hub.docker.com/settings/security
     - Create token with "Read, Write, Delete" permissions

### GitHub Container Registry
- ✅ `GITHUB_TOKEN` - Automatically provided by GitHub Actions (no setup needed)

## 🔧 Workflow Triggers

### Docker Workflow (`.github/workflows/docker.yml`)
- Triggers on:
  - Push to tags (e.g., `v1.0.0`)
  - Manual trigger (`workflow_dispatch`)
- Builds and pushes to:
  - Docker Hub: `zamibd/s-ui:latest`, `zamibd/s-ui:v1.0.0`, etc.
  - GHCR: `ghcr.io/zamibd/s-ui:latest`, `ghcr.io/zamibd/s-ui:v1.0.0`, etc.

### Release Workflow (`.github/workflows/release.yml`)
- Triggers on:
  - Push to `main` branch (when specific files change)
  - Release published
  - Manual trigger
- Builds Linux binaries for: amd64, arm64, armv7, armv6, armv5, 386, s390x
- Uploads to GitHub Releases

### Windows Workflow (`.github/workflows/windows.yml`)
- Triggers on:
  - Push to `main` branch (when specific files change)
  - Release published
  - Manual trigger
- Builds Windows binaries for: amd64, arm64
- Uploads to GitHub Releases

## 🚀 How to Deploy

### 1. First-time Setup
1. Configure GitHub Secrets (see above)
2. Push code to GitHub repository
3. Create and push a git tag to trigger Docker build:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

### 2. Manual Docker Build Trigger
- Go to GitHub Actions tab
- Select "Docker Image CI" workflow
- Click "Run workflow"
- Select branch and click "Run workflow"

### 3. Verify Deployment
- Docker Hub: https://hub.docker.com/r/zamibd/s-ui
- GHCR: https://github.com/zamibd/s-ui/pkgs/container/s-ui
- Check workflow runs in GitHub Actions tab

## 📝 Notes

1. **Docker Hub Repository**: Make sure the repository `zamibd/s-ui` exists on Docker Hub
2. **GHCR Visibility**: By default, GHCR packages are private. To make public:
   - Go to package settings
   - Change visibility to public
3. **Tag Naming**: The workflow uses semantic versioning tags (e.g., `v1.0.0`)
4. **Build Time**: Multi-platform Docker builds may take 10-20 minutes
5. **Cache**: Docker layer caching is enabled for faster subsequent builds

## 🔍 Testing Locally

Before pushing, test Docker build locally:
```bash
# Build frontend
cd frontend
npm install
npm run build
cd ..

# Build Docker image
docker build -f Dockerfile.frontend-artifact -t zamibd/s-ui:test .
docker run -p 2095:2095 -p 2096:2096 zamibd/s-ui:test
```

## ❌ Known Issues

None identified. All configurations appear correct.

