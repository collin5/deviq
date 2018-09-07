FROM python:slim

RUN pip install pytest

COPY ./startup.sh startup.sh

RUN chmod +x startup.sh

ENTRYPOINT [ "/bin/bash", "startup.sh" ]

CMD [ ]


