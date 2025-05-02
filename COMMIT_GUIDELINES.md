# Diretrizes para Mensagens de Commit

Estabelecemos regras precisas para como as mensagens de commit devem ser formatadas.  
Essas regras tornam as mensagens mais legíveis e facilitam a análise do histórico do projeto. 

---

## Formato da Mensagem de Commit

Cada mensagem de commit deve conter três partes principais:

```
<tipo>(<pacote>/<escopo>): <assunto>
```

Exemplo: 
feat(conversor): adiciona suporte a conversão para Kelvin

---

### Tipos de Commits

- **feat**: uma nova funcionalidade
- **fix**: uma correção de bug
- **docs**: mudanças apenas na documentação
- **style**: mudanças que não afetam a lógica (espaços, formatação, etc)
- **refactor**: refatorações de código
- **perf**: Melhorias de performance
- **test**: adição ou correção de testes
- **chore**: mudanças em tarefas de build ou ferramentas auxiliares
- **build**: Mudanças no sistema de build
- **ci**: Mudanças em configurações de integração contínua (CI)
- **release**: Marca um ponto de release na aplicação
- **revert**: Reversão de um commit anterior

---

## 📦 Pacote

Refere-se ao módulo, funcionalidade ou subcomponente impactado. Exemplos:

- `conversor`
- `input`
- `api`

---

## 🎯 Escopo

Define o local ou função específica da modificação dentro do pacote.  
Exemplos:

- `conversor/input`
- `api/service`
- `ui/botao`

---

## 📝 Assunto (Subject)

Deve ser uma descrição **curta e objetiva** da mudança.

Regras:
- Use o **imperativo no presente**: “corrige”, “adiciona”, “melhora”
- Não comece com letra maiúscula
- Mantenha a descrição curta (~72 caracteres).
- Se necessário, adicione uma descrição mais detalhada no corpo do commit.
- **Não** coloque ponto final

---

## ✅ Checklist

- [ ] Use o tipo correto (`feat`, `fix`, etc)
- [ ] Identifique o pacote e o escopo, se necessário
- [ ] Use o tempo verbal correto no assunto e corpo
- [ ] Evite linhas com mais de 100 caracteres
- [ ] Adicione links ou referências a issues se aplicável

---

Seguir esse padrão ajuda toda a equipe a entender o que foi feito, por que foi feito e onde foi feito com mais clareza.
