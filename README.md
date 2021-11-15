# check_local_memory

Plugin de supervision de la mémoire sur Linux pour Nagios/Icinga/Shinken/Centreon.

## Pour construire le binaire :

    # go build -o check_local_memory 

## Utilisation :

    $./check_local_memory -warning 80 -critical 90
