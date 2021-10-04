# Codanalyzer
コードを分析するアプリのドメインモデル

# コード設定
スキーマ設定

```yaml
CodeInputs:
    description: コードの入力値の集合
    type: object
    required: 
    - version
    - codes
    properties:
    codes:
      type: array
      description: コード構造の入力の集合
      items:
        type: object
        required: 
          - intervals
        properties:
          name:
           type: string
           description: コード名(オプション Em7b5やEdim7などの特殊なコードの場合に指定)
           example: 'M7'
          triad:
            type: string
            description: トライアド
            enum:
              - 'Major'
              - 'Minor'
              - 'Augment'
              - 'Diminish'
              - 'MajorB5'
              - 'Sus2'
              - 'Sus4'
          tension:
            type: array
            description: テンション
            example: ['7', 'b9']
            items:
              type: string
              enum: 
                - 'b6'
                - '6'
                - 'b7'
                - '7'
                - "b9"
                - "9"
                - "#9"
                - "11"
                - "#11"
                - "b13"
                - "13"
    
```