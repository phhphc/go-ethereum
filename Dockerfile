FROM archlinux:latest

RUN pacman -Sy --noconfirm geth
RUN echo 'geth --http --http.corsdomain="*" --http.api web3,eth,debug,personal,net --vmdebug --datadir /dev-chain --dev' >> /chain.sh

CMD ["/bin/sh", "/chain.sh"]