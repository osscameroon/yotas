docker run --name postgresql_bitnami \
    -e POSTGRESQL_DATABASE=yotas \
    -e POSTGRESQL_PASSWORD=root \
    -e POSTGRESQL_USERNAME=root \
    bitnami/postgresql:latest
