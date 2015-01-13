FROM dietfs
CMD rm -rf /bin/* /usr/bin/ /var/ /root /usr/share
ADD serve /bin/sh
WORKDIR /bin/
CMD /bin/sh
