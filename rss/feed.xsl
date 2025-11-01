<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
  <xsl:output method="html" encoding="UTF-8" indent="yes"/>
  
  <xsl:template match="/rss">
    <html lang="en">
      <head>
        <meta charset="UTF-8"/>
        <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
        <title><xsl:value-of select="channel/title"/></title>
        <style>
          * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
          }
          
          body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            background: #f5f5f5;
            padding: 20px;
          }
          
          .container {
            max-width: 900px;
            margin: 0 auto;
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
            overflow: hidden;
          }
          
          header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 40px;
            text-align: center;
          }
          
          header h1 {
            font-size: 2rem;
            font-weight: 600;
            margin-bottom: 10px;
          }
          
          header p {
            opacity: 0.9;
            font-size: 1rem;
          }
          
          .info-banner {
            background: #e8f4fd;
            border-left: 4px solid #2196F3;
            padding: 20px;
            margin: 20px 40px;
            border-radius: 4px;
          }
          
          .info-banner p {
            margin: 5px 0;
            font-size: 0.95rem;
            color: #1976D2;
          }
          
          .info-banner strong {
            font-weight: 600;
          }
          
          .content {
            padding: 40px;
          }
          
          .menu-items {
            display: grid;
            gap: 20px;
          }
          
          .menu-item {
            background: #fafafa;
            border: 1px solid #e0e0e0;
            border-radius: 8px;
            padding: 24px;
            transition: transform 0.2s, box-shadow 0.2s;
          }
          
          .menu-item:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
          }
          
          .menu-item h3 {
            font-size: 1.25rem;
            font-weight: 600;
            color: #333;
            margin-bottom: 12px;
            line-height: 1.4;
          }
          
          .menu-item .description {
            color: #666;
            font-size: 0.95rem;
            line-height: 1.5;
          }
          
          .menu-item .description i {
            font-style: normal;
            color: #2196F3;
            font-weight: 500;
          }
          
          footer {
            background: #f5f5f5;
            padding: 30px 40px;
            text-align: center;
            border-top: 1px solid #e0e0e0;
            font-size: 0.9rem;
            color: #666;
          }
          
          footer a {
            color: #667eea;
            text-decoration: none;
          }
          
          footer a:hover {
            text-decoration: underline;
          }
          
          .rss-icon {
            display: inline-block;
            width: 20px;
            height: 20px;
            background: #ff6600;
            border-radius: 3px;
            margin-right: 8px;
            vertical-align: middle;
          }
          
          @media (max-width: 768px) {
            body {
              padding: 10px;
            }
            
            header {
              padding: 30px 20px;
            }
            
            header h1 {
              font-size: 1.5rem;
            }
            
            .content {
              padding: 20px;
            }
            
            .info-banner {
              margin: 20px 20px;
              padding: 15px;
            }
            
            .menu-item {
              padding: 20px;
            }
            
            .menu-item h3 {
              font-size: 1.1rem;
            }
            
            footer {
              padding: 20px;
            }
          }
          
          @media (max-width: 480px) {
            header h1 {
              font-size: 1.25rem;
            }
            
            .menu-item {
              padding: 16px;
            }
            
            .menu-item h3 {
              font-size: 1rem;
            }
          }
        </style>
      </head>
      <body>
        <div class="container">
          <header>
            <h1>
              <span class="rss-icon"></span>
              <xsl:value-of select="channel/title"/>
            </h1>
            <p><xsl:value-of select="channel/description"/></p>
          </header>
          
          <div class="info-banner">
            <p><strong>📰 This is an RSS feed.</strong> Subscribe by copying the URL into your RSS reader.</p>
            <xsl:if test="channel/pubDate">
              <p><strong>Last updated:</strong> <xsl:value-of select="channel/pubDate"/></p>
            </xsl:if>
          </div>
          
          <div class="content">
            <div class="menu-items">
              <xsl:for-each select="channel/item">
                <div class="menu-item">
                  <h3><xsl:value-of select="title"/></h3>
                  <div class="description">
                    <!-- disable-output-escaping is safe here because the RSS feed is generated
                         by our own code which properly escapes HTML entities (e.g., &lt;br /&gt;).
                         We need to unescape them to render the HTML formatting. -->
                    <xsl:value-of select="description" disable-output-escaping="yes"/>
                  </div>
                </div>
              </xsl:for-each>
            </div>
          </div>
          
          <footer>
            <p>
              Generated by <xsl:value-of select="channel/generator"/> |
              <a href="{channel/link}" target="_blank">Visit Website</a>
            </p>
          </footer>
        </div>
      </body>
    </html>
  </xsl:template>
</xsl:stylesheet>
