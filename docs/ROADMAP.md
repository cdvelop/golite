## 📌 Hoja de Ruta

### ✅ MVP (Versión Mínima Viable)  
### Frontend
- [x] Unificación y minificación de archivos CSS y JavaScript 
- [ ] cargar assets del directorio `src/web/ui` primero (assets handler)
- [ ] Generación automática de `src/web/public/index.html` si este no existe  
- [ ] Compilar iconos svg módulos en sprite único en `src/web/public/icons.svg`

### Servidor de Desarrollo
- [ ] Servidor de desarrollo integrado para servir archivos estáticos en `src/web/public`
- [ ] https integrado en desarrollo local
- [x] cerrar navegador al cerrar aplicación 
- [x] Ejecución navegador Chrome (tecla `w`)  
- [x] cambiar el tamaño de la ventana del navegador desde la tui

### Hot Reload
- [x] Detección de cambios en archivos HTML, CSS, y JS (en `src/web/public` y `src/web/ui`)  
- [x] detección de cambios en archivos GO frontend para webAssembly (`src/cmd/webclient/`) y servidor backend (`src/cmd/appserver/`)
- [ ] detectar cambios en archivos SVG (en `src/web/ui`)
- [ ] Recarga en caliente del navegador (Hot Reload)

### Backend
- [x] Detección de cambios en archivos del servidor (`src/cmd/appserver/`)  
- [ ] Reinicio automático si hay modificaciones  

### Configuración
- [x] Interfaz TUI minimalista para VS Code  
- [x] **Detección automática por estructura de directorios** ✅
- [x] **Eliminación completa de archivos de configuración** ✅
- [ ] Finalizar especificación de interacción TUI
- [ ] Agregar .gitignore automático


### 🚀 Mejoras Futuras  
- [ ] **Completar especificación TUI** para interacción final
- [ ] Modo producción: Artefactos optimizados y despliegue
- [ ] Compatibilidad con servidores VPS
- [ ] Compatibilidad con Docker  
- [ ] Integrar MCP


**Instalador web automático** que detectará e instalará todo automáticamente:
- [ ] **Detección automática** de dependencias instaladas
- [ ] **Instalación automática** de faltantes: Go, Git, TinyGo, Docker, GitHub CLI  
- [ ] **Setup completo** con un solo comando
- [ ] **Sin conocimiento técnico** requerido