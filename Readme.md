# Strategy

At first glance, the problem can be almost entirely solved with a webserver like nginx.

However, implementing it in nginx will not allow us to showcase building / deploying docker images. Therefore I will implement this entirely in a programming language.

# Build and Deploy

scripts/build-and-push.sh | xargs scripts/cd.sh
