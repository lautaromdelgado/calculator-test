# 🧮 Calculator Test - Go Edition

¡Bienvenido a **Calculator Test**!\
Este proyecto es una **calculadora simple escrita en Go**, con foco en la
**división segura**, el **manejo de errores personalizados** y una
**arquitectura limpia y testeable**. Incluye **tests unitarios** utilizando el
poderoso framework [Testify](https://github.com/stretchr/testify).

---

## 📁 Estructura del Proyecto

```plaintext
calculator-test/
├── cmd/
│   └── main.go                 # Punto de entrada principal
├── internal/
│   ├── Calculator/
│   │   ├── divider.go         # Lógica de división
│   │   └── divider_test.go    # Pruebas unitarias
│   └── errors/
│       └── errors.go          # Definición de errores personalizados
├── go.mod                     # Archivo de definición de módulos
└── go.sum                     # Checksum de dependencias
```

---

## ⚙️ Funcionalidades Clave

- ✅ **División segura:**\
  La función `Divide` evita divisiones por cero y retorna un error específico
  (`ErrDivisionByZero`).

- 🧩 **Arquitectura modular:**\
  Uso de **interfaces** y separación por paquetes (`internal/Calculator`,
  `internal/errors`) para favorecer el **escalado** y **testing**.

- 🚨 **Manejo de errores personalizado:**\
  Todos los errores están definidos de forma centralizada en un paquete
  dedicado, mejorando la claridad y el control de excepciones.

---

## 🧪 Testing

Este proyecto utiliza
[testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite) y
[testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert) para
pruebas más robustas y legibles.

**Ventajas del uso de Testify:**

- Agrupación de tests en **Suites**.
- Métodos `SetupTest()` y `TearDownTest()` para preparar el entorno de pruebas
  automáticamente.
- **Aserciones claras**, fáciles de leer y mantener.

### ▶️ Ejecutar pruebas

```bash
go test ./internal/Calculator/...
```

---

## 💻 Ejemplo de Uso

```go
divider := calculator.NewDivider()

result, err := divider.Divide(10, 2)
// result = 5.0, err = nil

result, err = divider.Divide(10, 0)
// result = 0.0, err = ErrDivisionByZero
```

---

## 📌 Requisitos

- 🐹 **Go 1.18** o superior
- 📦 Conexión a internet para instalar dependencias

---

## 📥 Instalación de Dependencias

Ejecuta el siguiente comando en la raíz del proyecto:

```bash
go mod tidy
```

Esto descargará todas las dependencias necesarias para correr y testear la
aplicación.

---

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas!\
Sentite libre de crear issues, forks o pull requests si querés mejorar esta
calculadora o agregar nuevas funcionalidades.

---

## 🧑‍💻 Autor

Desarrollado por **Lautaro**, apasionado del backend con Go y defensor de las
buenas prácticas en código limpio y testeable.

---
