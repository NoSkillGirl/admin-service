FROM ubuntu:16.04
COPY admin-service /admin-service
RUN chmod +x /admin-service
CMD /admin-service
EXPOSE 8084