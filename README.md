# OpenMensaRSS

[![OpenMensaRSS Logo](rss/omrss.gif)](https://schicho.github.io/openmensarss/)

OpenMensa RSS library and feed generation in Go.

## Libary Usage

This library offers two simple functions to query OpenMensa using [go-openmensa](https://github.com/j0hax/go-openmensa) for a canteen, which return a [`gorilla/feeds` Feed](https://github.com/gorilla/feeds).

The returned feed structs can subsequently be modified and exported as RSS or Atom.

## How does this work?

Furthermore, this repository also implements automatic RSS feed generation. The daily generated feeds are then published on GitHub pages.
The RSS feed can then be used, embedded or read on any website or reader. For instance, you can add a RSS feed to your Tuwel frontpage.

The published RSS feeds use the same IDs as OpenMensa, e.g. the TU Wien canteen is served at https://schicho.github.io/openmensarss/1098.xml.

A GitHub Actions runner queries the OpenMensa API every day for selected canteens.
Using this glue code library the runner creates a RSS feed for each canteen.
Subsequently, the generated RSS feeds are deployed on GitHub pages.

More feeds can be found at https://schicho.github.io/openmensarss/.
