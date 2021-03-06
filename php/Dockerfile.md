```
FROM php:7.1.18-cli-stretch
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \
    && php -r "if (hash_file('SHA384', 'composer-setup.php') === '544e09ee996cdf60ece3804abc52599c22b1f40f4323403c44d44fdfdd586475ca9813a858088ffbc1f233e9b180f061') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;" \
    && php composer-setup.php \
    && php -r "unlink('composer-setup.php');" \
    && mv ./composer.phar /usr/local/bin/composer

RUN apt-get update

RUN apt-get install git --assume-yes \
    && apt-get install zlib1g-dev --assume-yes \
    && docker-php-source extract \
    && cd /usr/src/php/ext/zip \
    && phpize \
    && ./configure --with-php-config=/usr/local/bin/php-config \
    && make \
    && make install \
    && touch /usr/local/etc/php/conf.d/php-ext-zip.ini \
    && echo extension=zip.so >> /usr/local/etc/php/conf.d/php-ext-zip.ini \
    && docker-php-source delete

```