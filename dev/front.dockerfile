FROM node:20

RUN mkdir /app && chown node:node /app
WORKDIR /app

USER node

COPY --chown=node:node ../manage-console/package.json ../manage-console/yarn.lock ./

RUN yarn install

COPY --chown=node:node . .

EXPOSE 80

# CMD [ "yarn", "dev" ]