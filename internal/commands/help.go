package commands

import (
	"strings"
)

// Help devuelve un listado legible de todos los endpoints disponibles.
func Help() (string, error) {
	lines := []string{
		"/fibonacci?num={N}                 : Calcula el N-ésimo número de Fibonacci",
		"/createfile?name={name}&content={text}&repeat={times} : Crea o trunca un archivo con contenido repetido",
		"/deletefile?name={name}           : Elimina un archivo existente",
		"/reverse?text={text}              : Invierte la cadena de texto dada",
		"/toupper?text={text}              : Convierte el texto a mayúsculas",
		"/random?count={c}&min={min}&max={max} : Genera un arreglo de números aleatorios",
		"/timestamp                        : Devuelve la hora actual en formato ISO-8601",
		"/hash?text={text}                 : Calcula el hash SHA-256 del texto",
		"/simulate?seconds={s}&task={name} : Simula una tarea durmiendo X segundos",
		"/sleep?seconds={s}                : Suspende la ejecución durante X segundos",
		"/loadtest?tasks={n}&sleep={s}     : Ejecuta N tareas concurrentes durmiendo S segundos cada una",
		"/status                           : Muestra métricas del servidor (uptime, conexiones, procesos)",
		"/help                             : Muestra este mensaje de ayuda",
	}
	return strings.Join(lines, "\r\n"), nil
}
