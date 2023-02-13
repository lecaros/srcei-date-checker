# Buscador horas en Registro Civil

Me tocó varias veces buscar citas en la web del SRCEI que lamentablemente no tiene un mejor sistems de búsqueda si estás 
considerando ir a cualquier oficina mientras puedas hacer el trámite pronto.

Este finde fue lo mismo, pero esta vez me puse manos a la obra y escribí algo de Golang para automatizar la búsqueda de
próximas citas disponibles dentro de una región y para un trámite en particular.

> **Nota**: Este código está ~~muy~~ un poco sucio, pero funciona. Si alguien quiere hacer un PR para mejorarlo, 
> será bienvenido.

> **Nota 2**: Este código no está afiliado de ninguna manera con el SRCEI. Es solo un script que hice para mi uso personal.

> **Nota 3**: Usé Golang porque estoy aprendiendo y quería probarlo.


## Limitaciones

- Hay citas que aparecen disponibles, incluso en la web del SRCEI, pero luego no están disponibles. Esto es 
responsabilidad del SRCEI.
- No soporta múltiples trámites a la vez.
- No soporta múltiples regiones a la vez.



## Uso

```bash
$ go run main.go -t 2-5 -r 13
```

Esto buscará citas disponibles para el trámite de "Primera obtención de Cédula de Identidad para extranjero" (2-5) en las
oficinas de la RM (13).

Ejemplo de output
```bash
➜  srcei-date-checker go run ./main.go -t 2-5 -r 13
Buscando la cita más cercana para el trámite 2-5 en la región 13...
Oficina: Alhue - Fecha: 2023-03-16 10:10:00 +0000 UTC
Oficina: Atencion Especial Moneda - Fecha: 2023-04-03 10:48:00 +0000 UTC
Oficina: Buin - Fecha: 2023-05-05 11:40:00 +0000 UTC
Oficina: Calera De Tango - Fecha: 2023-04-13 09:50:00 +0000 UTC
Oficina: Cerro Navia - Fecha: 2023-04-06 13:30:00 +0000 UTC
Oficina: Colina - Fecha: 2023-04-25 10:06:00 +0000 UTC
Oficina: Conchali - Fecha: 2023-04-04 11:26:00 +0000 UTC
Oficina: Curacavi - Fecha: 2023-03-29 12:40:00 +0000 UTC
Oficina: El Bosque - Fecha: 2023-05-12 11:18:00 +0000 UTC
Oficina: El Monte - Fecha: 2023-04-05 12:40:00 +0000 UTC
Oficina: Estacion Central - Fecha: 2023-04-05 13:26:00 +0000 UTC
Oficina: Huechuraba - Fecha: 2023-04-03 10:38:00 +0000 UTC
Oficina: Huérfanos Piso 1 - Fecha: 2023-03-31 11:50:00 +0000 UTC
Oficina: Independencia - Fecha: 2023-04-10 11:42:00 +0000 UTC
Oficina: Isla De Maipo - Fecha: 2023-04-10 11:40:00 +0000 UTC
Oficina: La Cisterna - Fecha: 2023-04-12 10:38:00 +0000 UTC
Oficina: La Granja - Fecha: 2023-04-10 13:34:00 +0000 UTC
Oficina: La Pintana - Fecha: 2023-04-06 11:58:00 +0000 UTC
Oficina: La Reina - Fecha: 2023-04-10 10:46:00 +0000 UTC
Oficina: Lampa - Fecha: 2023-02-17 10:10:00 +0000 UTC
Oficina: Las Condes - Fecha: 2023-04-18 09:10:00 +0000 UTC
Oficina: Lo Barnechea - Fecha: 2023-04-13 11:58:00 +0000 UTC
Oficina: Lo Espejo - Fecha: 2023-05-03 10:20:00 +0000 UTC
Oficina: Lo Prado - Fecha: 2023-04-06 10:46:00 +0000 UTC
Oficina: Los Cerrillos - Fecha: 2023-04-21 09:18:00 +0000 UTC
Oficina: Macul - Fecha: 2023-04-18 12:14:00 +0000 UTC
Oficina: Maipu - Fecha: 2023-04-17 10:30:00 +0000 UTC
Oficina: Maria Pinto - Fecha: 2023-03-23 11:30:00 +0000 UTC
Oficina: Melipilla - Fecha: 2023-03-29 13:00:00 +0000 UTC
Oficina: Padre Hurtado - Fecha: 2023-05-12 10:20:00 +0000 UTC
Oficina: Paine - Fecha: 2023-05-19 10:10:00 +0000 UTC
Oficina: Pedro Aguirre Cerda - Fecha: 2023-04-03 11:18:00 +0000 UTC
Oficina: Peñaflor - Fecha: 2023-05-12 12:30:00 +0000 UTC
Oficina: Peñalolen - Fecha: 2023-04-19 12:30:00 +0000 UTC
Oficina: Pirque - Fecha: 2023-04-19 10:30:00 +0000 UTC
Oficina: Plaza Tobalaba - Fecha: 2023-04-24 11:18:00 +0000 UTC
Oficina: Providencia - Fecha: 2023-04-14 09:42:00 +0000 UTC
Oficina: Pudahuel - Fecha: 2023-04-20 09:34:00 +0000 UTC
Oficina: Puente Alto - Fecha: 2023-04-20 12:06:00 +0000 UTC
Oficina: Quilicura - Fecha: 2023-04-12 10:54:00 +0000 UTC
Oficina: Quinta Normal - Fecha: 2023-04-10 13:50:00 +0000 UTC
Oficina: Recoleta - Fecha: 2023-03-03 10:06:00 +0000 UTC
Oficina: Renca - Fecha: 2023-04-14 13:42:00 +0000 UTC
Oficina: San Bernardo - Fecha: 2023-04-17 12:22:00 +0000 UTC
Oficina: San Bernardo So Mall Plaza Sur - Fecha: 2023-04-17 09:26:00 +0000 UTC
Oficina: San Joaquin - Fecha: 2023-04-06 10:30:00 +0000 UTC
Oficina: San Jose De Maipo - Fecha: 2023-04-03 10:40:00 +0000 UTC
Oficina: San Miguel - Fecha: 2023-03-22 10:46:00 +0000 UTC
Oficina: San Pedro De Melipilla - Fecha: 2023-03-21 11:30:00 +0000 UTC
Oficina: San Ramón - Fecha: 2023-04-04 10:50:00 +0000 UTC
Oficina: So La Florida - Fecha: 2023-04-17 08:54:00 +0000 UTC
Oficina: Suboficina Maipu Arauco Maipu - Fecha: 2023-04-13 09:42:00 +0000 UTC
Oficina: Talagante - Fecha: 2023-04-11 11:26:00 +0000 UTC
Oficina: Til-Til - Fecha: 2023-03-27 13:20:00 +0000 UTC
Oficina: Vitacura - Fecha: 2023-03-01 09:18:00 +0000 UTC
Oficina: Ñuñoa - Fecha: 2023-04-24 10:46:00 +0000 UTC
===================================== Fechas más cercanas ordenadas. =====================================
Oficina: Lampa - Fecha: 2023-02-17 10:10:00 +0000 UTC
Oficina: Vitacura - Fecha: 2023-03-01 09:18:00 +0000 UTC
Oficina: Recoleta - Fecha: 2023-03-03 10:06:00 +0000 UTC
Oficina: Alhue - Fecha: 2023-03-16 10:10:00 +0000 UTC
Oficina: San Pedro De Melipilla - Fecha: 2023-03-21 11:30:00 +0000 UTC
Oficina: San Miguel - Fecha: 2023-03-22 10:46:00 +0000 UTC
Oficina: Maria Pinto - Fecha: 2023-03-23 11:30:00 +0000 UTC
Oficina: Til-Til - Fecha: 2023-03-27 13:20:00 +0000 UTC
Oficina: Curacavi - Fecha: 2023-03-29 12:40:00 +0000 UTC
Oficina: Melipilla - Fecha: 2023-03-29 13:00:00 +0000 UTC
Oficina: Huérfanos Piso 1 - Fecha: 2023-03-31 11:50:00 +0000 UTC
Oficina: Huechuraba - Fecha: 2023-04-03 10:38:00 +0000 UTC
Oficina: San Jose De Maipo - Fecha: 2023-04-03 10:40:00 +0000 UTC
Oficina: Atencion Especial Moneda - Fecha: 2023-04-03 10:48:00 +0000 UTC
Oficina: Pedro Aguirre Cerda - Fecha: 2023-04-03 11:18:00 +0000 UTC
Oficina: San Ramón - Fecha: 2023-04-04 10:50:00 +0000 UTC
Oficina: Conchali - Fecha: 2023-04-04 11:26:00 +0000 UTC
Oficina: El Monte - Fecha: 2023-04-05 12:40:00 +0000 UTC
Oficina: Estacion Central - Fecha: 2023-04-05 13:26:00 +0000 UTC
Oficina: San Joaquin - Fecha: 2023-04-06 10:30:00 +0000 UTC
Oficina: Lo Prado - Fecha: 2023-04-06 10:46:00 +0000 UTC
Oficina: La Pintana - Fecha: 2023-04-06 11:58:00 +0000 UTC
Oficina: Cerro Navia - Fecha: 2023-04-06 13:30:00 +0000 UTC
Oficina: La Reina - Fecha: 2023-04-10 10:46:00 +0000 UTC
Oficina: Isla De Maipo - Fecha: 2023-04-10 11:40:00 +0000 UTC
Oficina: Independencia - Fecha: 2023-04-10 11:42:00 +0000 UTC
Oficina: La Granja - Fecha: 2023-04-10 13:34:00 +0000 UTC
Oficina: Quinta Normal - Fecha: 2023-04-10 13:50:00 +0000 UTC
Oficina: Talagante - Fecha: 2023-04-11 11:26:00 +0000 UTC
Oficina: La Cisterna - Fecha: 2023-04-12 10:38:00 +0000 UTC
Oficina: Quilicura - Fecha: 2023-04-12 10:54:00 +0000 UTC
Oficina: Suboficina Maipu Arauco Maipu - Fecha: 2023-04-13 09:42:00 +0000 UTC
Oficina: Calera De Tango - Fecha: 2023-04-13 09:50:00 +0000 UTC
Oficina: Lo Barnechea - Fecha: 2023-04-13 11:58:00 +0000 UTC
Oficina: Providencia - Fecha: 2023-04-14 09:42:00 +0000 UTC
Oficina: Renca - Fecha: 2023-04-14 13:42:00 +0000 UTC
Oficina: So La Florida - Fecha: 2023-04-17 08:54:00 +0000 UTC
Oficina: San Bernardo So Mall Plaza Sur - Fecha: 2023-04-17 09:26:00 +0000 UTC
Oficina: Maipu - Fecha: 2023-04-17 10:30:00 +0000 UTC
Oficina: San Bernardo - Fecha: 2023-04-17 12:22:00 +0000 UTC
Oficina: Las Condes - Fecha: 2023-04-18 09:10:00 +0000 UTC
Oficina: Macul - Fecha: 2023-04-18 12:14:00 +0000 UTC
Oficina: Pirque - Fecha: 2023-04-19 10:30:00 +0000 UTC
Oficina: Peñalolen - Fecha: 2023-04-19 12:30:00 +0000 UTC
Oficina: Pudahuel - Fecha: 2023-04-20 09:34:00 +0000 UTC
Oficina: Puente Alto - Fecha: 2023-04-20 12:06:00 +0000 UTC
Oficina: Los Cerrillos - Fecha: 2023-04-21 09:18:00 +0000 UTC
Oficina: Ñuñoa - Fecha: 2023-04-24 10:46:00 +0000 UTC
Oficina: Plaza Tobalaba - Fecha: 2023-04-24 11:18:00 +0000 UTC
Oficina: Colina - Fecha: 2023-04-25 10:06:00 +0000 UTC
Oficina: Lo Espejo - Fecha: 2023-05-03 10:20:00 +0000 UTC
Oficina: Buin - Fecha: 2023-05-05 11:40:00 +0000 UTC
Oficina: Padre Hurtado - Fecha: 2023-05-12 10:20:00 +0000 UTC
Oficina: El Bosque - Fecha: 2023-05-12 11:18:00 +0000 UTC
Oficina: Peñaflor - Fecha: 2023-05-12 12:30:00 +0000 UTC
Oficina: Paine - Fecha: 2023-05-19 10:10:00 +0000 UTC
La mejor fecha es: 2023-02-17 10:10:00 +0000 UTC en la oficina Lampa

```

## ToDo

Se reciben PRs y sugerencias en los issues.