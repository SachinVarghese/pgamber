FROM postgres:14.5

RUN apt-get update 
RUN apt-get -y install python3 postgresql-plpython3-14 python3-pip
RUN pip install alibi-detect[tensorflow]

ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 5432
CMD ["postgres"]