#! /usr/bin/zsh

cur_dir="$(dirname "$(test -L "$0" && readlink "$0" || echo "$0")")"
url="https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=$MAX_MIND_KEY&suffix=tar.gz"
curl $url -o $cur_dir/tmp.tar.gz
tar -xzf $cur_dir/tmp.tar.gz --wildcards --no-anchored -C $cur_dir '*.mmdb'
mv $cur_dir/Geo*/*.mmdb $cur_dir/geo.mmdb
rm -rf $cur_dir/tmp.tar.gz $cur_dir/Geo*
mv $cur_dir/geo.mmdb $cur_dir/../assets
