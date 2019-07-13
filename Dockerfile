FROM scratch
ENV PORT 5000
COPY hellohttp /hellohttp
EXPOSE 5000/tcp
CMD ["/hellohttp"]