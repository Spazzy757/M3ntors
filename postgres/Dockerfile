FROM postgres:14

RUN apt update && \
  apt install -y curl  && \
  apt-get clean autoclean && \
  apt-get autoremove --yes && \
  rm -rf /var/lib/{apt,dpkg,cache,log}
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash -
RUN apt-get install -y nodejs
RUN mkdir -p /home/postgres/.nvm
RUN mkdir /home/postgres/migrations
RUN chown -R postgres:postgres /home/postgres

WORKDIR /home/postgres/migrations

COPY ./scripts/entrypoint.sh /docker-entrypoint-initdb.d/

USER postgres
