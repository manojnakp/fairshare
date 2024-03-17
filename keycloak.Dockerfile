FROM registry.access.redhat.com/ubi9:9.3 AS ubi-micro-build
RUN mkdir -p /mnt/rootfs
RUN dnf install --installroot /mnt/rootfs wget --releasever 9 --setopt install_weak_deps=false --nodocs -y && \
    dnf --installroot /mnt/rootfs clean all && \
    rpm --root /mnt/rootfs -e --nodeps setup

FROM quay.io/keycloak/keycloak:24.0
COPY --from=ubi-micro-build /mnt/rootfs /
ENV KC_HEALTH_ENABLED=true
EXPOSE 8080
ENTRYPOINT ["/opt/keycloak/bin/kc.sh"]
CMD ["start-dev"]
