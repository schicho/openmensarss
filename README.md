# OpenMensaRSS

OpenMensa RSS library and feed.

## How does this work?

A GitHub Actions runner queries the OpenMensa API every day for selected canteens.
Using this glue code library the runner creates a RSS feed for each canteen.
Subsequently, the generated RSS feeds are deployed on GitHub pages at
https://schicho.github.io/openmensarss/.

The RSS feed can then be used, embedded or read on any website or reader.
For instance, you can add a RSS feed to your Tuwel frontpage.
