#! /bin/bash

configFile=$1
path_to_migration_files=$2
force_version=$3
PORT=$(jq -r '.DATABASE.DB_PORT' $configFile)
USERNAME=$(jq -r '.DATABASE.DB_USERNAME' $configFile)
DBNAME=$(jq -r '.DATABASE.DB_NAME' $configFile)
PASSWORD=$(jq -r '.DATABASE.DB_PASSWORD' $configFile)
POSTGRESQL_URL=postgres://${USERNAME}:${PASSWORD}@localhost:${PORT}/${DBNAME}?sslmode=disable

echo 'Runing migrations...'
echo ''

PS3='Please enter your choice: '
options=("migration_up" "migration_down" "migration_force" "Quit")
echo ''

select opt in "${options[@]}"
do
    case $opt in
        "migration_up")
            echo 'Runing migrations UP...'
            $(migrate -database ${POSTGRESQL_URL} -path $path_to_migration_files up)
            ;;
        "migration_down")
            echo 'Runing migrations DOWN...'
            $(migrate -database ${POSTGRESQL_URL} -path $path_to_migration_files down)
            ;;
        "migration_force")
            echo 'Runing migrations FORCE VERSION...'
            $(migrate -path migrations/ -database ${POSTGRESQL_URL} -path $path_to_migration_files force $force_version)
            ;;

        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done


